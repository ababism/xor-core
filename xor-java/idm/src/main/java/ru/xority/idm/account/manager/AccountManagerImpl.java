package ru.xority.idm.account.manager;

import java.util.UUID;

import org.springframework.stereotype.Service;

import ru.xority.idm.account.dao.AccountDao;
import ru.xority.idm.account.pojo.Account;

/**
 * @author foxleren
 */
@Service
public class AccountManagerImpl implements AccountManager {
    private final AccountDao accountDao;

    public AccountManagerImpl(AccountDao accountDao) {
        this.accountDao = accountDao;
    }

    @Override
    public void register(String login, String password) {
        // TODO hash password
        String passwordHash = password;
        accountDao.create(Account.getRegisterAccount(login, passwordHash));
    }

    @Override
    public void updatePassword(UUID id, String password) {
        // TODO hash password
        String passwordHash = password;
        accountDao.updatePassword(id, passwordHash);
    }
}
