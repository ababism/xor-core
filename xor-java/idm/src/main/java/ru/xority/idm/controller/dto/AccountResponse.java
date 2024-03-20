package ru.xority.idm.controller.dto;

import java.util.Optional;
import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@JsonInclude(JsonInclude.Include.NON_EMPTY)
public class AccountResponse {
    @JsonProperty("uuid")
    private UUID uuid;
    @JsonProperty("email")
    private String email;
    @JsonProperty("active")
    private boolean active;
    @JsonProperty("firstName")
    private Optional<String> firstName;
    @JsonProperty("lastName")
    private Optional<String> lastName;
    @JsonProperty("telegramUsername")
    private Optional<String> telegramUsername;
}
