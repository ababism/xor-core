package ru.xority.idm.service;

import java.util.List;
import java.util.Optional;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import ru.xority.idm.entity.AccountEntity;
import ru.xority.idm.entity.AccountFilter;
import ru.xority.idm.repository.AccountRepository;

/**
 * @author foxleren
 */
@Service
@RequiredArgsConstructor
public class AccountService {
    private final AccountRepository accountRepository;

    public List<AccountEntity> list(AccountFilter filter) {
        return accountRepository.list(filter);
    }

    public Optional<AccountEntity> get(AccountFilter filter) {
        return accountRepository.get(filter);
    }
}
