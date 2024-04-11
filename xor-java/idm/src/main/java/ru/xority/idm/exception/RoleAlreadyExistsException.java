package ru.xority.idm.exception;

import ru.xority.exception.BadRequestException;

/**
 * @author foxleren
 */
public class RoleAlreadyExistsException extends BadRequestException {
    public RoleAlreadyExistsException() {
        super("Role already exists");
    }
}
