const fonts = [
  ["Andale Mono", "mono"],
  ["Arial Black", "sans"],
  ["Arial Hebrew", "sans"],
  ["Arial MT", "sans"],
  ["Arial Narrow", "sans"],
  ["Arial Rounded MT Bold", "sans"],
  ["Arial Unicode MS", "sans"],
  ["Arial", "sans"],
  ["Bitstream Vera Sans Mono", "mono"],
  ["Book Antiqua", "serif"],
  ["Bookman Old Style", "serif"],
  ["Calibri", "sans"],
  ["Cambria", "serif"],
  ["Century Gothic", "serif"],
  ["Century Schoolbook", "serif"],
  ["Century", "serif"],
  ["Comic Sans MS", "sans"],
  ["Comic Sans", "sans"],
  ["Consolas", "mono"],
  ["Courier New", "mono"],
  ["Courier", "mono"],
  ["Garamond", "serif"],
  ["Georgia", "serif"],
  ["Helvetica Neue", "sans"],
  ["Helvetica", "sans"],
  ["Impact", "sans"],
  ["Lucida Fax", "serif"],
  ["Lucida Handwriting", "script"],
  ["Lucida Sans Typewriter", "mono"],
  ["Lucida Sans Unicode", "sans"],
  ["Lucida Sans", "sans"],
  ["MS Gothic", "sans"],
  ["MS Outlook", "symbol"],
  ["MS PGothic", "sans"],
  ["MS Reference Sans Serif", "sans"],
  ["MS Serif", "serif"],
  ["MYRIAD PRO", "sans"],
  ["MYRIAD", "sans"],
  ["Microsoft Sans Serif", "sans"],
  ["Monaco", "sans"],
  ["Monotype Corsiva", "script"],
  ["Palatino Linotype", "serif"],
  ["Palatino", "serif"],
  ["Segoe Script", "script"],
  ["Segoe UI Semibold", "sans"],
  ["Segoe UI Symbol", "symbol"],
  ["Segoe UI", "sans"],
  ["Tahoma", "sans"],
  ["Times New Roman PS", "serif"],
  ["Times New Roman", "serif"],
  ["Times", "serif"],
  ["Trebuchet MS", "sans"],
  ["Verdana", "sans"],
  ["Wingdings 3", "symbol"],
  ["Wingdings", "symbol"],
];

const baseFonts = ["monospace", "sans-serif", "serif"];

const ctx = document.createElement("canvas").getContext("2d");

const safeMeasureText = (fonts) => {
  ctx.font = "72px " + fonts.join(", ");
  try {
    return ctx.measureText("mmmmmmmmmmlli").width;
  } catch (e) {
    return -1;
  }
};

const fontFp = () => {
  return new Promise((res) => {
    const start = performance.now();
    const baseFontsFp = {};

    baseFonts.forEach((font) => {
      baseFontsFp[font] = safeMeasureText([font, "monospace"]);
    });

    let fp = "";
    let elapsed = performance.now() - start;
    let i = 0;
    const t = () => {
      if (i >= fonts.length) {
        return res({ value: fp, calculationTimeMs: elapsed });
      }
      const font = fonts[i];
      const baseFont = "monospace" === font[1] ? "sans-serif" : "monospace";
      const measurement = safeMeasureText([font[0], baseFont]);
      fp += baseFontsFp[baseFont] !== measurement ? "1" : "0";
      elapsed += performance.now() - start;
      i++;
      setTimeout(t);
    };

    t();
  });
};
