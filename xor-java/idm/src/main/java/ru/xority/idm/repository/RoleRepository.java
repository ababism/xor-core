package ru.xority.idm.repository;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import ru.xority.idm.entity.RoleEntity;
import ru.xority.idm.entity.RoleFilter;

/**
 * @author foxleren
 */
public interface RoleRepository {
    List<RoleEntity> list(RoleFilter filter);

    Optional<RoleEntity> get(RoleFilter filter);

    void create(RoleEntity role);

    void update(RoleEntity role);

    void assignRole(UUID accountUuid, UUID roleUuid);

    void revoke(UUID accountUuid, UUID roleUuid);
}
