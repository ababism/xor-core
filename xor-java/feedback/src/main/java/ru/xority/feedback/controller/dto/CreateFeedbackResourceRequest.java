package ru.xority.feedback.controller.dto;

import jakarta.validation.constraints.NotBlank;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
public class CreateFeedbackResourceRequest {
    @NotBlank
    private String name;
    @NotBlank
    private String description;
}
