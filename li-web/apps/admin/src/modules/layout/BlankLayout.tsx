export function BlankLayout(props: any) {
  return (
    <div
      style={{
        maxWidth: 320,
        margin: "0 auto",
        paddingTop: "20vh",
      }}
    >
      <h1>Li</h1>
      {props.children}
    </div>
  );
}
