package ru.xority.idm.controller.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class AccessTokenResponse {
    @JsonProperty("accessToken")
    private String accessToken;
}
