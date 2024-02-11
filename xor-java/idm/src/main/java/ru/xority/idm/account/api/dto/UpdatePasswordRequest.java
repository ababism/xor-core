package ru.xority.idm.account.api.dto;

import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * @author foxleren
 */
public record UpdatePasswordRequest(
        @JsonProperty("id")
        UUID id,
        @JsonProperty("password")
        String password)
{
}
