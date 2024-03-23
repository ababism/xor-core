package ru.xority.feedback.controller.dto;

import java.util.UUID;

import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Positive;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
public class UpdateFeedbackItemInfoRequest {
    @NotNull
    private UUID itemUuid;
    @NotBlank
    private String text;
    @Positive
    @Min(1)
    @Max(5)
    private int rating;
}
