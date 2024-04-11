package ru.xority.feedback.controller.dto;

import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;

/**
 * @author foxleren
 */
@Data
@AllArgsConstructor
public class GetFeedbackResourceStatsResponse {
    @JsonProperty("uuid")
    private UUID uuid;
    @JsonProperty("average_rating")
    private double averageRating;
}
