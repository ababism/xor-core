package ru.xority.utils;

import java.util.List;
import java.util.Optional;

/**
 * @author foxleren
 */
public class DataFilter {
    public static <T> T single(List<T> items) {
        if (items.size() != 1) {
            throw new RuntimeException("List must contain single result");
        }
        return items.get(0);
    }

    public static <T> Optional<T> singleO(List<T> items) {
        if (items.size() != 1) {
            return Optional.empty();
        }
        return Optional.of(items.get(0));
    }
}
