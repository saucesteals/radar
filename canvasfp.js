const canvasFp = () => {
  const canvas = document.createElement("canvas");
  try {
    canvas.height = 60;
    canvas.width = 400;
    canvas.style.display = "inline";
    var ctx = canvas.getContext("2d");
    ctx.textBaseline = "alphabetic";
    ctx.fillStyle = "#f60";
    ctx.fillRect(125, 1, 62, 20);
    ctx.fillStyle = "#069";
    ctx.font = "11pt Arial";
    ctx.fillText("Cwm fjordbank glyphs vext quiz, 😃", 2, 15);
    ctx.fillStyle = "rgba(102, 204, 0, 0.7)";
    ctx.font = "18pt Arial";
    ctx.fillText("Cwm fjordbank glyphs vext quiz, 😃", 4, 45);
    return md5(canvas.toDataURL());
  } catch (e) {
    return "unavailable";
  }
};
