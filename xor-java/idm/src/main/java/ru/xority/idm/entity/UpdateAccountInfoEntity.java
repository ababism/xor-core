package ru.xority.idm.entity;

import java.util.Optional;

import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class UpdateAccountInfoEntity {
    private Optional<String> firstName;
    private Optional<String> lastName;
    private Optional<String> telegramUsername;
}
