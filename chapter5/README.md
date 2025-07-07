

# 実装
### button1
ボタン1を押したらLEDが消えるが重い処理を走らせるとその処理を待つのでラグが発生する

### button2
button1の改善。
SetInterruptにボタン1を押した時の挙動を書いて割り込み処理をしている

## 用語

- **ペリフェラル**: マイコンに内蔵された機能。

- **マイクロコントローラ**: マイコン。マイクロコントローラ（MCUまたはマイコン）は、CPU、メモリ（RAM、ROM/フラッシュメモリ）、I/Oポート、タイマーなどの周辺機能を1つのチップにまとめた小型のコンピュータです。家電製品、自動車、産業機器、IoTデバイスなど、私たちの身の回りのあらゆる電子機器の制御に使われています。

- **マイコンシリーズ**: 「マイクロコントローラシリーズ」とは、ある半導体メーカーが提供する、共通のアーキテクチャや特徴を持つマイクロコントローラ（マイコン）の製品群を指します。しばしば「ファミリー」とも呼ばれます。

- **ATSAMD51**: Wio Terminalで使われているマイコン。さまざまなペリフェラルを持つ。

- **12分解能**
アナログ信号（光や音から計算した電圧など）をデジタル値に変換する際、そのアナログ値の範囲を2の12乗=4096段階のデジタル値で表現できるということ。つまりアナログ信号の細かな変化を4096段階でデジタル値に変換できる。(ATSAMD51はこれだが他のマイコンはわからない)

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



## minicomの設定

minicomで使うポートとwioterminalで使うポートを揃える必要がある？
以下で設定を開く



```
minicom -s
```

```

            +-----[configuration]------+
            | Filenames and paths      |
            | File transfer protocols  |
            | Serial port setup        |
            | Modem and dialing        |
            | Screen                   |
            | Keyboard and Misc        |
            | Save setup as dfl        |
            | Save setup as..          |
            | Exit                     |
            | Exit from Minicom        |
            +--------------------------+
```
Serial port setupを選択

```
    +-----------------------------------------------------------------------+
    | A -    Serial Device      : /dev/modem                                |
    | B - Lockfile Location     : /usr/local/Cellar/minicom/2.10/var        |
    | C -   Callin Program      :                                           |
    | D -  Callout Program      :                                           |
    | E -    Bps/Par/Bits       : 115200 8N1                                |
    | F - Hardware Flow Control : No                                        |
    | G - Software Flow Control : No                                        |
    | H -     RS485 Enable      : No                                        |
    | I -   RS485 Rts On Send   : No                                        |
    | J -  RS485 Rts After Send : No                                        |
    | K -  RS485 Rx During Tx   : No                                        |
    | L -  RS485 Terminate Bus  : No                                        |
    | M - RS485 Delay Rts Before: 0                                         |
    | N - RS485 Delay Rts After : 0                                         |
    |                                                                       |
    |    Change which setting?                                              |
    +-----------------------------------------------------------------------+
```

A Serial Deviceが `/dev/modem`になっているのでここを直してみる


---

# PWM

## TCC
Timer/Counter for Controlの略

## デューティー比

パルスが「オン（HIGH）」になっている時間の、周期全体に対する割合のこと。
明るさや音量などはこの仕組みで制御されている
![Image](https://github.com/user-attachments/assets/48ae139d-1c70-412b-926a-c252b7c351df)

ディーティー比が高い = オン(High)になっている割合が高い = モニタが明るい

## PWM
Pluse Width Modulation

段階調節をして出力したい時に使う
machine.LEDだとON/OFFだけしか制御できないがpwmを使うと薄く点灯したりできる

## チャンネル
チャンネル = マイコン内部のPWM出力回路 ピン = 外部デバイスとの接続点

```
マイコン内部:
TCC0 (PWMコントローラ) 
├── Channel A ──→ 物理ピン1 ──→ BUZZER_CTR
├── Channel B ──→ 物理ピン2 ──→ LED_PIN  
├── Channel C ──→ 物理ピン3 ──→ MOTOR_PIN
└── Channel D ──→ 物理ピン4 ──→ SERVO_PIN
```

・・・なぜチャンネルが必要？
ハードウェアの制約
1つのPWMコントローラで複数のピンを同時制御
各ピンは異なるチャンネルに割り当てられている
同じ周波数でもデューティ比は個別設定可能

例えば以下のようにブザーを鳴らしながらLEDを光らせるような時にチャンネルを使い分ける

```
// より実用的な例: ブザー + LED + モーター
channelA, _ := pwm.Channel(machine.BUZZER_CTR) // 内蔵ブザー
channelB, _ := pwm.Channel(machine.LED_PIN)    // LED明度制御
channelC, _ := pwm.Channel(machine.MOTOR_PIN)  // モーター速度制御

pwm.SetPeriod(1000) // 1kHz
pwm.Set(channelA, pwm.Top()/2)  // ブザー: 50%
pwm.Set(channelB, pwm.Top()/4)  // LED: 25%明度  
pwm.Set(channelC, pwm.Top()*3/4) // モーター: 75%速度
```