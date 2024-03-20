package ru.xority.idm.repository;

import java.util.List;
import java.util.Optional;

import ru.xority.idm.entity.AccountEntity;
import ru.xority.idm.entity.AccountFilter;

/**
 * @author foxleren
 */
public interface AccountRepository {
    List<AccountEntity> list(AccountFilter filter);

    Optional<AccountEntity> get(AccountFilter filter);

    void create(AccountEntity account);

    void update(AccountEntity account);
}
