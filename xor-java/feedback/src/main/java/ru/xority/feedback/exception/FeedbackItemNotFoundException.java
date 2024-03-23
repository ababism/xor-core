package ru.xority.feedback.exception;

import ru.xority.exception.BadRequestException;

/**
 * @author foxleren
 */
public class FeedbackItemNotFoundException extends BadRequestException {
    public FeedbackItemNotFoundException() {
        super("Feedback item is not found");
    }
}
