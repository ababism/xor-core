package ru.xority.idm.exception;

import ru.xority.exception.BadRequestException;

/**
 * @author foxleren
 */
public class RoleNotFoundException extends BadRequestException {
    public RoleNotFoundException() {
        super("Role is not found");
    }
}
