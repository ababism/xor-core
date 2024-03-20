package ru.xority.response;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class BadRequestResponse {
    @JsonProperty("message")
    private String message;

    public static ResponseEntity<BadRequestResponse> create(HttpStatus httpStatus, String message) {
        return new ResponseEntity<>(new BadRequestResponse(message), httpStatus);
    }
}
