components {
  id: "enemyBullet"
  component: "/Assets/scripts/enemyBullet.script"
}
embedded_components {
  id: "bullet"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/builtins/graphics/particle_blob.tilesource\"\n"
  "}\n"
  ""
  scale {
    x: 0.5
    y: 0.5
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 0.1\n"
  "friction: 0.0\n"
  "restitution: 0.5\n"
  "group: \"playerBullet\"\n"
  "mask: \"ground\"\n"
  "mask: \"enemy\"\n"
  "mask: \"boxes\"\n"
  "mask: \"walls\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 5.0\n"
  "}\n"
  "bullet: true\n"
  ""
}
