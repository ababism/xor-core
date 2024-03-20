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
public class AccountFilter {
    private Optional<UUID> uuid;
    private Optional<String> email;
    private Optional<Boolean> active;

    public static AccountFilter byUuid(UUID uuid) {
        return new AccountFilter(
                Optional.of(uuid),
                Optional.empty(),
                Optional.empty()
        );
    }

    public static AccountFilter activeByEmail(String email) {
        return new AccountFilter(
                Optional.empty(),
                Optional.of(email),
                Optional.of(true)
        );
    }
}
