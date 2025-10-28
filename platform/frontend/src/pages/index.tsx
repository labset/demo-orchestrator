import { BrowserRouter, Navigate, Route, Routes } from "react-router";
import { HomePage } from "./home-page";

const Pages = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/home" element={<HomePage />} />
        <Route path="*" element={<Navigate to="/home" replace />} />
      </Routes>
    </BrowserRouter>
  );
};

export { Pages };
