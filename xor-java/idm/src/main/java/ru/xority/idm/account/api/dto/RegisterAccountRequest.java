package ru.xority.idm.account.api.dto;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * @author foxleren
 */
public record RegisterAccountRequest(
        @JsonProperty("login")
        String login,
        @JsonProperty("password")
        String password)
{
}
