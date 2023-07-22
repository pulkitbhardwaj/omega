import { Button, Card, Text, TextInput } from "@tremor/react";
import Sticky from "react-stickynode";

/* eslint-disable-next-line */
export interface HeaderProps {}

export function Header(props: HeaderProps) {
  return (
    <Sticky className="fixed top-0">
      <Card
        className="w-screen h-24 rounded-none flex justify-between align-middle"
        decoration="bottom"
        decorationColor="indigo"
      >
        {/* <Flex> */}
        {/* Hamburger menu */}
        <Button variant="light">
          <div className="space-y-2">
            <div className="w-8 h-0.5 bg-gray-400"></div>
            <div className="w-8 h-0.5 bg-gray-400"></div>
            <div className="w-8 h-0.5 bg-gray-400"></div>
          </div>
        </Button>
        <TextInput placeholder="Search..." className="w-1/2" />
        <Text>Logo</Text>
        {/* </Flex> */}
      </Card>
    </Sticky>
  );
}

export default Header;
