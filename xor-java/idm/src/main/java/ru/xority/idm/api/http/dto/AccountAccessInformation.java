package ru.xority.idm.api.http.dto;

import java.util.List;
import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class AccountAccessInformation {
    @JsonProperty("uuid")
    private UUID uuid;
    @JsonProperty("email")
    private String email;
    @JsonProperty("roles")
    private List<String> roles;
}
