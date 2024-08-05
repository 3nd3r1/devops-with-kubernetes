const imagenatorUrl =
  process.env.NEXT_PUBLIC_IMAGENATOR_URL ?? "/imagenator/image";

const ImagenatorImage = () => {
  return <img src={imagenatorUrl} alt="Imagenator" />;
};

export default ImagenatorImage;
