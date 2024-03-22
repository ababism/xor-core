package ru.xority.notifications.controller.dto;

import java.util.UUID;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
public class CreatePlatformNotificationRequest {
    @NotNull
    private UUID receiverUuid;
    @NotBlank
    private String message;
}
