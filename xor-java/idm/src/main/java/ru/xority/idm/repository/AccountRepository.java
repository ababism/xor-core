package ru.xority.idm.repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import ru.xority.idm.entity.AccountEntity;
import ru.xority.idm.entity.AccountFilter;
import ru.xority.idm.entity.RoleEntity;

/**
 * @author foxleren
 */
public interface AccountRepository {
    List<AccountEntity> list(AccountFilter filter);

    Optional<AccountEntity> get(AccountFilter filter);

    void create(AccountEntity account);

    void update(AccountEntity account);

    List<RoleEntity> getActiveRoles(UUID accountUuid);


//    List<AccountRole> getRoles(UUID accountUuid);
//
//    void assignRole(AccountRole role);
}
