import { Spin } from "@arco-design/web-react";

export const Loading: React.FC = () => {
  return (
    <div
      style={{
        width: "100vw",
        height: "100vh",
        position: "relative",
      }}
    >
      <Spin
        size={100}
        tip="Loading..."
        style={{
          position: "absolute",
          left: "50%",
          top: "50%",
          transform: "translateX(-50%) translateY(-50%)",
        }}
      />
    </div>
  );
};
