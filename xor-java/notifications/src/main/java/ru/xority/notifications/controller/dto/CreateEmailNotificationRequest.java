package ru.xority.notifications.controller.dto;

import jakarta.validation.constraints.NotBlank;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
public class CreateEmailNotificationRequest {
    @NotBlank
    private String receiverEmail;
    @NotBlank
    private String subject;
    @NotBlank
    private String body;
}
