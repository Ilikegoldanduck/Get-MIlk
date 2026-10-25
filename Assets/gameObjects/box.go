components {
  id: "box"
  component: "/Assets/scripts/box.script"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.1\n"
  "group: \"boxes\"\n"
  "mask: \"ground\"\n"
  "mask: \"player\"\n"
  "mask: \"playerBullet\"\n"
  "mask: \"enemy\"\n"
  "mask: \"enemyBullet\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 16.0\n"
  "  data: 16.0\n"
  "  data: 10.0\n"
  "}\n"
  "locked_rotation: true\n"
  ""
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"boxes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 32.0\n"
  "  y: 32.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/Assets/imgbin/boxes_copy.tilesource\"\n"
  "}\n"
  ""
  scale {
    x: 2.0
    y: 2.0
  }
}
