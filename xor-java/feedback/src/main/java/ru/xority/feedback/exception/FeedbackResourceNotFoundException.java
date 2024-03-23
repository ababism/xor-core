package ru.xority.feedback.exception;

import ru.xority.exception.BadRequestException;

/**
 * @author foxleren
 */
public class FeedbackResourceNotFoundException extends BadRequestException {
    public FeedbackResourceNotFoundException() {
        super("Feedback resource is not found");
    }
}
