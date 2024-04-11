package ru.xority.idm.api.http;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestHeader;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import ru.xority.idm.api.http.dto.AssignRoleRequest;
import ru.xority.idm.api.http.dto.CreateRoleRequest;
import ru.xority.idm.api.http.dto.RevokeRoleRequest;
import ru.xority.idm.api.http.dto.RoleResponse;
import ru.xority.idm.entity.RoleEntity;
import ru.xority.idm.entity.RoleFilter;
import ru.xority.idm.service.RoleService;
import ru.xority.response.SuccessResponse;
import ru.xority.sage.SageHeader;

/**
 * @author foxleren
 */
@RestController
@RequestMapping("/admin/role")
@RequiredArgsConstructor
public class RoleController {
    private final RoleService roleService;

    @GetMapping("/list")
    public ResponseEntity<?> list(@RequestParam Optional<UUID> uuid,
                                  @RequestParam Optional<String> name) {
        RoleFilter filter = new RoleFilter(uuid, name);
        List<RoleResponse> roles = roleService
                .list(filter)
                .stream()
                .map(RoleResponse::fromRoleEntity)
                .toList();
        return ResponseEntity.ok(roles);
    }

    @PostMapping("/create")
    public ResponseEntity<?> create(@RequestBody @Valid CreateRoleRequest request,
                                    @RequestHeader(SageHeader.XOR_ACCOUNT_UUID) UUID createdByUuid) {
        RoleEntity role = RoleEntity.fromCreateRequest(request.getName(), createdByUuid);
        roleService.create(role);
        return SuccessResponse.create200("Role is created");
    }

    @PutMapping("/set-active/{uuid}")
    public ResponseEntity<?> deactivate(@PathVariable UUID uuid,
                                        @RequestParam boolean active) {
        roleService.setActive(uuid, active);
        return SuccessResponse.create200("Role active status is updated");
    }

    @PostMapping("/assign")
    public ResponseEntity<?> assign(@RequestBody @Valid AssignRoleRequest request) {
        roleService.assign(request.getAccountUuid(), request.getRoleUuid());
        return SuccessResponse.create200("Role is assigned");
    }

    @PostMapping("/revoke")
    public ResponseEntity<?> revoke(@RequestBody @Valid RevokeRoleRequest request) {
        roleService.revoke(request.getAccountUuid(), request.getRoleUuid());
        return SuccessResponse.create200("Role is revoked");
    }
}
