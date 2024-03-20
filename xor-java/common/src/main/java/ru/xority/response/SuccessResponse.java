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
public class SuccessResponse {
    @JsonProperty("message")
    private String message;

    public static ResponseEntity<SuccessResponse> create200(String message) {
        return new ResponseEntity<>(new SuccessResponse(message), HttpStatus.OK);
    }
}
