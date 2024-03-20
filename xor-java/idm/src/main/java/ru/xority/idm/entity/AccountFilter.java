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
    private boolean active;

    public static AccountFilter byEmail(String email) {
        return new AccountFilter(
                Optional.empty(),
                Optional.of(email),
                true
        );
    }
}
