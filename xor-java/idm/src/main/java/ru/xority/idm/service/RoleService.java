package ru.xority.idm.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import ru.xority.exception.BadRequestException;
import ru.xority.idm.entity.RoleEntity;
import ru.xority.idm.entity.RoleFilter;
import ru.xority.idm.exception.RoleAlreadyExistsException;
import ru.xority.idm.exception.RoleNotFoundException;
import ru.xority.idm.repository.AccountRepository;
import ru.xority.idm.repository.RoleRepository;

/**
 * @author foxleren
 */
@Service
@RequiredArgsConstructor
public class RoleService {
    private final RoleRepository roleRepository;
    private final AccountRepository accountRepository;

    public List<RoleEntity> list(RoleFilter filter) {
        return roleRepository.list(filter);
    }

    public void create(RoleEntity role) {
        RoleFilter filter = new RoleFilter(
                Optional.empty(),
                Optional.of(role.getName())
        );
        Optional<RoleEntity> roleO = roleRepository.get(filter);
        if (roleO.isPresent()) {
            throw new RoleAlreadyExistsException();
        }

        roleRepository.create(role);
    }

    public void setActive(UUID uuid, boolean active) {
        Optional<RoleEntity> roleO = roleRepository.get(RoleFilter.byUuid(uuid));
        if (roleO.isEmpty()) {
            throw new RoleNotFoundException();
        }
        RoleEntity role = roleO.get();
        if (role.isActive() == active) {
            throw new BadRequestException("Role active status is not changed");
        }

        role.setActive(active);
        roleRepository.update(role);
    }

    public void assign(UUID accountUuid, UUID roleUuid) {
        Optional<RoleEntity> roleO = roleRepository.get(RoleFilter.byUuid(roleUuid));
        if (roleO.isEmpty()) {
            throw new RoleNotFoundException();
        }
        RoleEntity role = roleO.get();
        if (!role.isActive()) {
            throw new RoleNotFoundException();
        }

        // TODO проверить, если такая роль уже есть
        roleRepository.assignRole(accountUuid, roleUuid);
    }

    public void revoke(UUID accountUuid, UUID roleUuid) {
        roleRepository.revoke(accountUuid, roleUuid);
    }
}
