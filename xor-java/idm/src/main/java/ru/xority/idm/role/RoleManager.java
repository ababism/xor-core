package ru.xority.idm.role;

/**
 * @author foxleren
 */
public class RoleManager {
    private final RoleDao roleDao;

    public RoleManager(RoleDao roleDao) {
        this.roleDao = roleDao;
    }
}
