package ru.xority.feedback.controller.dto;

import java.util.UUID;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
public class UpdateFeedbackResourceInfoRequest {
    @NotNull
    private UUID resourceUuid;
    @NotBlank
    private String name;
    @NotBlank
    private String description;
}
