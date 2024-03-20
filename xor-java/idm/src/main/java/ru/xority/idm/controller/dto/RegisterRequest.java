package ru.xority.idm.controller.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
public class RegisterRequest {
    @JsonProperty("email")
    @NotBlank
    @Email
    private String email;
    @JsonProperty("password")
    @NotBlank
    private String password;
}
