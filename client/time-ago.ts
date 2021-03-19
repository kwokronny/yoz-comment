const thresholds = [1000, "秒", 1000 * 60, "分", 1000 * 60 * 60, "时", 1000 * 60 * 60 * 24, "天", 1000 * 60 * 60 * 24 * 7, "周", 1000 * 60 * 60 * 24 * 27, "月"];

const formatOptions: Intl.DateTimeFormatOptions = { month: "short", day: "numeric", year: "numeric" };

export function timeAgo(value: string) {
	let date:Date = new Date(value);
  const elapsed = new Date().getTime() - new Date(value).getTime();
  if (elapsed < 5000) {
    return "刚刚";
  }
  let i = 0;
  while (i + 2 < thresholds.length && elapsed * 1.1 > thresholds[i + 2]) {
    i += 2;
  }

  const divisor = thresholds[i] as number;
  const text = thresholds[i + 1] as string;
  const units = Math.round(elapsed / divisor);

  if (units > 3 && i === thresholds.length - 2) {
    return date.toLocaleDateString(undefined, formatOptions);
  }
  return `${units}${text}前`;
}
