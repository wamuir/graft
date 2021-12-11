package(
    default_visibility = ["//:__subpackages__"],
)

constraint_setting(
    name = "device",
    default_constraint_value = ":cpu",
)

constraint_value(
    name = "cpu",
    constraint_setting = ":device",
)

constraint_value(
    name = "gpu",
    constraint_setting = ":device",
)
