package ru.xority.notifications.controller.dto;

import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class CreateEmailNotificationResponse {
    private UUID uuid;
}
