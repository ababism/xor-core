package ru.xority.exception;

/**
 * @author foxleren
 */
public class BadRequestException extends IllegalArgumentException {
    public BadRequestException(String message) {
        super(message);
    }
}
