package ru.xority.notifications.controller.dto;

import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class CreatePlatformNotificationResponse {
    @JsonProperty("uuid")
    private UUID uuid;
}
