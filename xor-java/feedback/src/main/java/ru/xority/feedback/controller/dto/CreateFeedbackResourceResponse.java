package ru.xority.feedback.controller.dto;

import java.util.UUID;

import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class CreateFeedbackResourceResponse {
    private UUID uuid;
}
