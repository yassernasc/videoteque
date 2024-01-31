export const PlayerLegacy = () => (
  <div>
    <video
      // eslint-disable-next-line tailwindcss/no-custom-classname
      className="video-js mx-auto block"
      controls
      preload="auto"
      width="720"
      height="480"
      data-setup="{}"
    >
      <source src="/movie" type="video/mp4"></source>
      <track src="/subtitle" label="Subtitles" default />
    </video>
  </div>
)
