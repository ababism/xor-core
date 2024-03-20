package ru.xority.idm.exception;

import ru.xority.exception.BadRequestException;

/**
 * @author foxleren
 */
public class AccountNotFoundException extends BadRequestException {
    public AccountNotFoundException() {
        super("Account is not found");
    }
}
