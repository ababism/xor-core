package ru.xority.idm.entity;

import java.util.Optional;
import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class RoleFilter {
    private Optional<UUID> uuid;
    private Optional<String> name;

    public static RoleFilter byUuid(UUID uuid) {
        return new RoleFilter(
                Optional.of(uuid),
                Optional.empty()
        );
    }
}
