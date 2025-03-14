# ImgTool - CLI Image Resizer & Converter

**ImgTool** is a powerful command-line tool built with Go and Cobra that allows users to resize and convert images efficiently.

## 🚀 Features

- Resize images while maintaining aspect ratio
- Convert images to different formats (JPEG, PNG, WebP)
- Easy-to-use command-line interface
- Supports batch processing (upcoming feature)

---

## 📌 Installation

### 1️⃣ Install Go (if not installed)

Ensure you have Go installed. If not, download and install it from [here](https://go.dev/dl/).

### 2️⃣ Install Cobra CLI

```bash
go install github.com/spf13/cobra-cli@latest
```

### 3️⃣ Clone the repository & Build

```bash
git clone https://github.com/your-username/imgtool.git
cd imgtool
go build -o imgtool
```

### 4️⃣ Verify Installation

```bash
./imgtool --help
```

---

## 📖 Usage

### 🔹 Convert an Image

```bash
./imgtool convert --input sample.jpg --output output.png --format png
```

- `--input`  (or `-i`) : Path to the input image
- `--output` (or `-o`) : Path to the output image
- `--format` (or `-f`) : Output format (jpeg, png, webp)

### 🔹 Resize an Image

```bash
./imgtool resize --input sample.jpg --output resized.jpg --width 800 --height 600
```

- `--input`  (or `-i`) : Path to the input image
- `--output` (or `-o`) : Path to the output image
- `--width`  (or `-w`) : New width
- `--height` (or `-h`) : New height

---

## 🛠️ Dependencies

- **Go** (>=1.17)
- [Cobra CLI](https://github.com/spf13/cobra)
- [disintegration/imaging](https://github.com/disintegration/imaging) (for image manipulation)
- [chai2010/webp](https://github.com/chai2010/webp) (for WebP support)

---

## 📌 Roadmap

✅ Image conversion (JPEG, PNG, WebP)  
✅ Image resizing with aspect ratio  
🔜 Bat
