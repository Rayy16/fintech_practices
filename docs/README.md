# fintech_practices

## 音色克隆

```py

python gen_voice.py ./sample/<text_file.txt> ./sample/<your_wav_file.wav>

```

## 数字人合成

```py

python inference.py --driven_audio ./sample/<audio.wav> \
                    --source_image ./sample/<video.mp4 or picture.png> \
                    --result_dir ./res \
                    --enhancer gfpgan \
                    --still \
                    --preprocess full 

```