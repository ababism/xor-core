package ru.xority.idm.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import lombok.RequiredArgsConstructor;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import ru.xority.exception.BadRequestException;
import ru.xority.idm.entity.AccountEntity;
import ru.xority.idm.entity.AccountFilter;
import ru.xority.idm.entity.UpdateAccountInfoEntity;
import ru.xority.idm.exception.AccountNotFoundException;
import ru.xority.idm.repository.AccountRepository;

/**
 * @author foxleren
 */
@Service
@RequiredArgsConstructor
public class AccountService {
    private final AccountRepository accountRepository;
    private final PasswordEncoder passwordEncoder;

    public UserDetailsService userDetailsService() {
        return email -> accountRepository.get(AccountFilter.activeByEmail(email))
                .orElseThrow(() -> new UsernameNotFoundException("User not found"));
    }

    public List<AccountEntity> list(AccountFilter filter) {
        return accountRepository.list(filter);
    }

    public Optional<AccountEntity> get(AccountFilter filter) {
        return accountRepository.get(filter);
    }

    public void setActive(UUID uuid, boolean active) {
        Optional<AccountEntity> accountO = accountRepository.get(AccountFilter.byUuid(uuid));
        if (accountO.isEmpty()) {
            throw new AccountNotFoundException();
        }
        AccountEntity account = accountO.get();
        account.setActive(active);
        accountRepository.update(account);
    }

    public void updatePassword(UUID uuid, String password) {
        Optional<AccountEntity> accountO = accountRepository.get(AccountFilter.byUuid(uuid));
        if (accountO.isEmpty()) {
            throw new AccountNotFoundException();
        }
        AccountEntity account = accountO.get();
        if (passwordEncoder.matches(password, account.getPasswordHash())) {
            throw new BadRequestException("New account password matches old password");
        }
        account.setPasswordHash(passwordEncoder.encode(password));
        accountRepository.update(account);
    }

    public void updateInfo(UUID uuid, UpdateAccountInfoEntity entity) {
        Optional<AccountEntity> accountO = accountRepository.get(AccountFilter.byUuid(uuid));
        if (accountO.isEmpty()) {
            throw new AccountNotFoundException();
        }
        AccountEntity account = accountO.get();
        account.setFirstName(entity.getFirstName());
        account.setLastName(entity.getLastName());
        account.setTelegramUsername(entity.getTelegramUsername());
        accountRepository.update(account);
    }
}
