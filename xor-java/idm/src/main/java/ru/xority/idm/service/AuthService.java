package ru.xority.idm.service;

import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import ru.xority.exception.BadRequestException;
import ru.xority.idm.common.jwt.JwtService;
import ru.xority.idm.common.jwt.JwtTokenParams;
import ru.xority.idm.entity.AccountEntity;
import ru.xority.idm.entity.AccountFilter;
import ru.xority.idm.exception.AccountNotFoundException;
import ru.xority.idm.repository.AccountRepository;

/**
 * @author foxleren
 */
@Service
@RequiredArgsConstructor
public class AuthService {
    private final AccountRepository accountRepository;
    private final PasswordEncoder passwordEncoder;
    private final JwtService jwtService;

    public String register(String email, String password) {
        Optional<AccountEntity> accountFromRepo = accountRepository.get(AccountFilter.byEmail(email));
        if (accountFromRepo.isPresent()) {
            throw new BadRequestException("Email is already registered");
        }
        String passwordHash = passwordEncoder.encode(password);
        AccountEntity account = AccountEntity.createdAccount(email, passwordHash);
        accountRepository.create(account);

        JwtTokenParams jwtTokenParams = jwtTokenParamsFromAccountEntity(account);
        return jwtService.generateToken(jwtTokenParams);
    }

    public String logIn(String email, String password) {
        Optional<AccountEntity> accountO = accountRepository.get(AccountFilter.byEmail(email));
        if (accountO.isEmpty()) {
            throw new AccountNotFoundException();
        }

        AccountEntity account = accountO.get();
        if (!passwordEncoder.matches(password, account.getPasswordHash())) {
            throw new AccountNotFoundException();
        }

        JwtTokenParams jwtTokenParams = jwtTokenParamsFromAccountEntity(accountO.get());
        return jwtService.generateToken(jwtTokenParams);
    }

    private static JwtTokenParams jwtTokenParamsFromAccountEntity(AccountEntity account) {
        JwtTokenParams jwtTokenParams = new JwtTokenParams();
        jwtTokenParams.setSubject(account.getEmail());
        return jwtTokenParams;
    }
}
