package ru.xority.idm.api.http.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class AccessTokenResponse {
    @JsonProperty("access_token")
    private String accessToken;
}
