import { Group, Title } from "@mantine/core";
import { IconTilde } from "@tabler/icons-react";

const TopNavigation = () => {
  return (
    <Group h="100%" px="md" justify="space-between">
      <Group gap="xs">
        <Group gap={4}>
          <IconTilde size={20} stroke={1.5} />
          <Title order={4}>Product</Title>
        </Group>
      </Group>
    </Group>
  );
};

export { TopNavigation };
