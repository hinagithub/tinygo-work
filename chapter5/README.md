## 用語

- **ペリフェラル**: マイコンに内蔵された機能。

- **マイクロコントローラ**: マイコン。マイクロコントローラ（MCUまたはマイコン）は、CPU、メモリ（RAM、ROM/フラッシュメモリ）、I/Oポート、タイマーなどの周辺機能を1つのチップにまとめた小型のコンピュータです。家電製品、自動車、産業機器、IoTデバイスなど、私たちの身の回りのあらゆる電子機器の制御に使われています。

- **マイコンシリーズ**: 「マイクロコントローラシリーズ」とは、ある半導体メーカーが提供する、共通のアーキテクチャや特徴を持つマイクロコントローラ（マイコン）の製品群を指します。しばしば「ファミリー」とも呼ばれます。

- **ATSAMD51**: Wio Terminalで使われているマイコン。さまざまなペリフェラルを持つ。

---

## GPIOについて

Wio Terminalには、GPIO (General Purpose Input/Output) という、一般的な用途に使える電気の入出力ピンがたくさんあります。

例えるなら、**Wio Terminalの「手足」や「感覚器官」**のようなものです。

- **入力 (Input)**: ボタンが押されたか（ON/OFF）、センサーから光の強さ（値）が送られてきたか、などをWio Terminalが**「感じ取る」**ために使います。
- **出力 (Output)**: LEDを光らせる（ON/OFF）、モーターを回す（ON/OFF）、ブザーを鳴らす、などをWio Terminalが**「何かをする」**ために使います。

これらのGPIOピンを使うことで、Wio Terminalにさまざまな電子部品（センサー、モーター、LEDなど）を接続して、外部の世界とやり取りするプログラムを作ることができます。

---

## 出力

#### LED点灯回路例

この回路は、Wio Terminalのデジタル出力ピン（GPIOピン）である **D0ピン** を使って、LEDを点灯・消灯させるものです。

<img width="294" alt="Image" src="https://github.com/user-attachments/assets/4f0f3e2d-0b03-4fdc-a38d-b600d61da552" />

```
VCC3v3 → (制御するMCUのGPIOピン) → R52（電流制限） → USER LED（光る） → GND
```

---

#### USER LED回路の一般的な動作イメージ

- **VCC3v3**: 電源のプラス側（3.3V）。
- **マイクロコントローラ (MCU) のGPIOピン**: USER LEDは、通常、MCUの汎用入出力（GPIO）ピンに接続されています。MCUがこのGPIOピンを「HIGH」（3.3Vに近い状態）または「LOW」（0V/GNDに近い状態）に設定することでLEDのオン/オフを制御します。
- **R52（抵抗器）**: LEDを直列に接続する場合、電流制限抵抗として機能します。LEDは低い順方向電圧（例: 赤色LEDで約2V）で動作し、過剰な電流が流れると壊れてしまいます。R52は、LEDに流れる電流を適切な値に制限し、LEDを保護する役割を担います。
- **USER LED**: R52を通過した電流がLEDに流れ込み、LEDが発光します。
- **GND (グラウンド)**: 回路の負極（0V）。電流はLEDを通過した後、GNDに戻ります。

---

### P15

LEDは **P15番** に接続されています。

---

### 回路図記号の略号

| 略号 | 日本語名 (英語名) | 例 |
|------|-------------------|----|
| R    | 抵抗器 (Resistor) | R1, R51, RLOAD など |
| C    | コンデンサ (Capacitor) | C1, CBYPASS, CDECOUPLING など |
| L    | コイル、インダクタ (Inductor) | L1, LFILTER など |
| D    | ダイオード (Diode) | D1, DLED (LEDもダイオードの一種) |
| Q    | トランジスタ (Transistor) | Q1, QMOSFET など |
| U    | 集積回路 (Integrated Circuit - IC) | U1, UMCU (マイクロコントローラ), UADC |
| IC   | 集積回路 (Integrated Circuit - IC) | IC1, ICMCU など |
| J    | コネクタ、ジャック (Jack, Connector) | J1, JUART, JUSB など |
| P    | プラグ、ピンヘッダ (Plug, Pin Header) | P1, P15 (ピン番号を表すことも) |
| SW   | スイッチ (Switch) | SW1, SWRESET など |
| F    | ヒューズ (Fuse) | F1 |
| X    | 水晶振動子 (Crystal) | X1 |
| TP   | テストポイント (Test Point) | TP1 |
| VR   | 可変抵抗器 (Variable Resistor - ポテンショメータ) | VR1 |
| Z    | ツェナーダイオード、インピーダンス素子 | Z1 |



## 入力

なんでBUTTON3だけB始まりのアルファベットなんだろ？

<img width="708" alt="Image" src="https://github.com/user-attachments/assets/04de1cac-7b2a-44af-9df7-86fa2d21bf4a" />


> Buttonこれは、回路図やピン配置図でよく見られる、物理的なコネクタやバスのグループ分け、または特定の機能ブロックの区別を示している可能性が高いです。

