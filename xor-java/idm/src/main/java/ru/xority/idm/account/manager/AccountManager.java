package ru.xority.idm.account.manager;

import java.util.UUID;

/**
 * @author foxleren
 */
public interface AccountManager {
    void register(String login, String password);

    void updatePassword(UUID id, String password);
}
