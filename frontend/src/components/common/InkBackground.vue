<template>
  <canvas ref="canvasRef" class="ink-background"></canvas>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'

const canvasRef = ref<HTMLCanvasElement | null>(null)
let ctx: CanvasRenderingContext2D | null = null
let animationId: number | null = null

interface InkDrop {
  x: number
  y: number
  radius: number
  maxRadius: number
  opacity: number
  expansionRate: number
  fadeRate: number
  layers: InkLayer[]
}

interface InkLayer {
  radius: number
  opacity: number
  color: string
}

const inkDrops: InkDrop[] = []

const createInkDrop = (x?: number, y?: number) => {
  const canvas = canvasRef.value
  if (!canvas) return

  const maxRadius = Math.random() * 200 + 150
  const baseOpacity = Math.random() * 0.4 + 0.3

  const drop: InkDrop = {
    x: x ?? Math.random() * canvas.width,
    y: y ?? Math.random() * canvas.height,
    radius: 0,
    maxRadius,
    opacity: baseOpacity,
    expansionRate: Math.random() * 1.5 + 0.8,
    fadeRate: Math.random() * 0.003 + 0.001,
    layers: [
      { radius: 0, opacity: baseOpacity, color: '#1a1a2e' },
      { radius: 0, opacity: baseOpacity * 0.7, color: '#16213e' },
      { radius: 0, opacity: baseOpacity * 0.4, color: '#2c3e50' },
    ],
  }
  inkDrops.push(drop)
}

const drawInkDrop = (drop: InkDrop) => {
  if (!ctx) return

  drop.layers.forEach((layer, index) => {
    const layerRadius = drop.radius * (1 - index * 0.2)
    if (layerRadius <= 0) return

    const gradient = ctx!.createRadialGradient(
      drop.x,
      drop.y,
      0,
      drop.x,
      drop.y,
      layerRadius
    )

    const layerOpacity = drop.opacity * (1 - index * 0.3)
    gradient.addColorStop(0, hexToRgba(layer.color, layerOpacity))
    gradient.addColorStop(0.3, hexToRgba(layer.color, layerOpacity * 0.8))
    gradient.addColorStop(0.7, hexToRgba(layer.color, layerOpacity * 0.4))
    gradient.addColorStop(1, hexToRgba(layer.color, 0))

    ctx!.fillStyle = gradient
    ctx!.beginPath()
    ctx!.arc(drop.x, drop.y, layerRadius, 0, Math.PI * 2)
    ctx!.fill()
  })

  ctx.globalCompositeOperation = 'multiply'
}

const hexToRgba = (hex: string, alpha: number): string => {
  const r = parseInt(hex.slice(1, 3), 16)
  const g = parseInt(hex.slice(3, 5), 16)
  const b = parseInt(hex.slice(5, 7), 16)
  return `rgba(${r}, ${g}, ${b}, ${alpha})`
}

const animate = () => {
  const canvas = canvasRef.value
  if (!canvas || !ctx) return

  ctx.clearRect(0, 0, canvas.width, canvas.height)

  for (let i = inkDrops.length - 1; i >= 0; i--) {
    const drop = inkDrops[i]
    if (!drop) continue
    if (drop.radius < drop.maxRadius) {
      drop.radius += drop.expansionRate
    }
    
    drop.opacity -= drop.fadeRate

    if (drop.opacity <= 0) {
      inkDrops.splice(i, 1)
    } else {
      drawInkDrop(drop)
    }
  }

  ctx.globalCompositeOperation = 'source-over'

  animationId = requestAnimationFrame(animate)
}

const handleMouseMove = (e: MouseEvent) => {
  const canvas = canvasRef.value
  if (!canvas) return

  const rect = canvas.getBoundingClientRect()
  const x = e.clientX - rect.left
  const y = e.clientY - rect.top

  if (Math.random() < 0.3) {
    createInkDrop(x, y)
  }
}

const handleClick = (e: MouseEvent) => {
  const canvas = canvasRef.value
  if (!canvas) return

  const rect = canvas.getBoundingClientRect()
  const x = e.clientX - rect.left
  const y = e.clientY - rect.top

  for (let i = 0; i < 3; i++) {
    setTimeout(() => {
      createInkDrop(
        x + (Math.random() - 0.5) * 100,
        y + (Math.random() - 0.5) * 100
      )
    }, i * 100)
  }
}

const resizeCanvas = () => {
  const canvas = canvasRef.value
  if (!canvas) return

  canvas.width = window.innerWidth
  canvas.height = window.innerHeight
}

onMounted(() => {
  const canvas = canvasRef.value
  if (!canvas) return

  ctx = canvas.getContext('2d')
  resizeCanvas()

  window.addEventListener('resize', resizeCanvas)
  window.addEventListener('mousemove', handleMouseMove)
  window.addEventListener('click', handleClick)

  for (let i = 0; i < 5; i++) {
    setTimeout(() => createInkDrop(), i * 300)
  }

  animate()
})

onUnmounted(() => {
  if (animationId) {
    cancelAnimationFrame(animationId)
  }
  window.removeEventListener('resize', resizeCanvas)
  window.removeEventListener('mousemove', handleMouseMove)
  window.removeEventListener('click', handleClick)
})
</script>

<style scoped>
.ink-background {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: -1;
  background: linear-gradient(135deg, #f5f5f0 0%, #faf8f5 100%);
}
</style>
