import { Button, Card, Flex, Text, TextInput } from "@tremor/react";
import { Link } from "react-router-dom";
import Dashboard from "../dashboard/dashboard";

/* eslint-disable-next-line */
export interface LoginProps {}

export function Login(props: LoginProps) {
  return (
    <Flex
      className="h-screen w-screen bg-slate-500"
      justifyContent="center"
      alignItems="center"
    >
      <Card
        className="max-w-md mx-auto px-16 py-32"
        decoration="top"
        decorationColor="indigo"
      >
        <Flex flexDirection="col" justifyContent="between" alignItems="stretch">
          <TextInput placeholder="Email Address" className="mb-8" />
          <TextInput placeholder="Password" className="mb-4" />
          <Text className="mb-12">Forgot Password?</Text>
          <Link to="">
          <Button size="md">Login</Button>
          </Link>
        </Flex>
      </Card>
    </Flex>
  );
}

export default Login;
