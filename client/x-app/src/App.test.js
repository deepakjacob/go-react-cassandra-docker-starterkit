import React from "react";
import { mount } from "enzyme";
import App from "./App";
import Main from "./components/Main";
import MuiThemeProvider from "@material-ui/core/styles/MuiThemeProvider";
describe("App", () => {
  it("renders a <Main/> component", () => {
    expect.assertions(2);
    const wrapper = mount(<App />);
    expect(wrapper.contains(MuiThemeProvider)).toBe(true);
    expect(wrapper.contains(Main)).toBe(true);
  });
});
