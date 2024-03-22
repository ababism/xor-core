package ru.xority.response;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import ru.xority.exception.BadRequestException;

/**
 * @author foxleren
 */
@RestControllerAdvice
public class DefaultExceptionController {
    @ExceptionHandler(BadRequestException.class)
    public ResponseEntity<BadRequestResponse> handleBadRequestException(BadRequestException e) {
        return BadRequestResponse.create(HttpStatus.BAD_REQUEST, e.getMessage());
    }
}
