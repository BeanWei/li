/**
 * 图表图片导出, 只支持 renderer = 'canvas'
 * @param chart chart 实例
 * @param name 图片名称，可选，默认名为 'download'
 */
export function downloadImage(chart: any, name: string = "download") {
  if (!chart) return;
  const canvas = chart.getCanvas();
  canvas.get("timeline").stopAllAnimations();
  const dataURL = canvas.get("el").toDataURL({
    format: "png",
    quality: 1,
  });
  let a: HTMLAnchorElement | null = document.createElement("a");
  a.href = dataURL;
  a.download = `${name}-${new Date().toLocaleDateString()}.png`;
  a.click();
}
